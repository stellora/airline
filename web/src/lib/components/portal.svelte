<script module lang="ts">
	import { tick, type Snippet } from 'svelte'

	/**
	 * Usage: `<div use:portal={{target: "css selector", replace}}>` or `<div use:portal={{ target:
	 * document.body, replace }}>`.
	 */
	export function portal(
		el: HTMLElement,
		{ target, replace }: { target: HTMLElement | string; replace: boolean },
	) {
		let targetEl: Element | null | undefined
		let contentNodes: ChildNode[] | undefined
		let replaced: ChildNode[] | undefined
		async function update(newTarget: HTMLElement | string) {
			target = newTarget
			if (typeof target === 'string') {
				targetEl = document.querySelector(target)
				if (targetEl === null) {
					await tick()
					targetEl = document.querySelector(target)
				}
				if (!targetEl) {
					throw new Error(`No element found matching css selector: "${target}"`)
				}
			} else if (target instanceof HTMLElement) {
				targetEl = target
			} else {
				throw new TypeError(
					`Unknown portal target type: ${
						target === null ? 'null' : typeof target
					}. Allowed types: string (CSS selector) or HTMLElement.`,
				)
			}

			contentNodes = Array.from(el.childNodes)
			if (replace) {
				replaced = Array.from(targetEl.childNodes)
				for (const el of replaced) {
					if (el instanceof HTMLElement) {
						el.dataset.prevDisplay = el.style.display
						el.style.display = 'none'
					}
				}
			}
			targetEl.append(...contentNodes)
		}

		function destroy() {
			if (contentNodes) {
				for (const contentNode of contentNodes) {
					contentNode.remove()
				}
			}
			if (replaced) {
				for (const el of replaced) {
					if (el instanceof HTMLElement) {
						el.style.display = el.dataset.prevDisplay ?? 'block'
					}
				}
			}
		}

		update(target)
		return {
			update,
			destroy,
		}
	}
</script>

<script lang="ts">
	const {
		target,
		replace = false,
		children,
	}: { target: string | HTMLElement; replace?: boolean; children?: Snippet<[]> } = $props()
</script>

<div use:portal={{ target, replace }} hidden>
	{@render children?.()}
</div>
