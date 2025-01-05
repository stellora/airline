<script module lang="ts">
	import { tick, type Snippet } from 'svelte'

	/**
	 * Usage: `<div use:portal="css selector">` or `<div use:portal={document.body}>`
	 */
	export function portal(
		el: HTMLElement,
		{ target, replace }: { target: HTMLElement | string; replace: boolean },
	) {
		let targetEl
		async function update(newTarget: HTMLElement | string) {
			target = newTarget
			if (typeof target === 'string') {
				targetEl = document.querySelector(target)
				if (targetEl === null) {
					await tick()
					targetEl = document.querySelector(target)
				}
				if (targetEl === null) {
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
			if (replace) {
				targetEl.replaceChildren(el)
			} else {
				targetEl.appendChild(el)
			}
			el.hidden = false
		}

		function destroy() {
			if (el.parentNode) {
				el.parentNode.removeChild(el)
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
		replace = false,
	}: { target: string | HTMLElement; children?: Snippet<[]>; replace?: boolean } = $props()
</script>

<div use:portal={{ target, replace }} hidden>
	{@render children?.()}
</div>
