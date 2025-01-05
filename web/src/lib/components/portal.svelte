<script module lang="ts">
	import { tick, type Snippet } from 'svelte'

	/**
	 * Usage: `<div use:portal="css selector">` or `<div use:portal={document.body}>`.
	 */
	export function portal(el: HTMLElement, target: HTMLElement | string) {
		let targetEl: Element | null | undefined
		let contentNodes: ChildNode[] | undefined
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
			targetEl.append(...contentNodes)
		}

		function destroy() {
			if (contentNodes) {
				for (const contentNode of contentNodes) {
					contentNode.remove()
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
		children,
	}: {
		target: string | HTMLElement
		children?: Snippet<[]>
	} = $props()
</script>

<div use:portal={target} hidden>
	{@render children?.()}
</div>
