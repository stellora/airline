type Item = string
export type AdminBreadcrumbEntry = Item | Promise<Item> | Promise<AdminBreadcrumbEntry[]>

/**
 * Used in a `+page.server.ts` file's `load` function to provide a breadcrumb entry for a page.
 */
export async function breadcrumbEntry(
	parent: (() => Promise<{ breadcrumbs: AdminBreadcrumbEntry[] }>) | null,
	item: Item | Promise<Item>
): Promise<{ breadcrumbs: AdminBreadcrumbEntry[] }> {
	const parentBreadcrumbs = parent ? await parent().then(({ breadcrumbs }) => breadcrumbs) : null
	return { breadcrumbs: parentBreadcrumbs ? [...parentBreadcrumbs, item] : [item] }
}
