type Item = string
export type BreadcrumbEntry = Item | Promise<Item> | Promise<BreadcrumbEntry[]>

/**
 * Used in a `+page.server.ts` file's `load` function to provide a breadcrumb entry for a page.
 */
export async function breadcrumbEntry(
	parent: (() => Promise<{ breadcrumbs: BreadcrumbEntry[] }>) | null,
	item: Item | Promise<Item>
): Promise<{ breadcrumbs: BreadcrumbEntry[] }> {
	const parentBreadcrumbs = parent ? await parent().then(({ breadcrumbs }) => breadcrumbs) : null
	return { breadcrumbs: parentBreadcrumbs ? [...parentBreadcrumbs, item] : [item] }
}
