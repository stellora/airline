type Item = string
export type AdminBreadcrumbEntry = Item | Promise<Item> | Promise<AdminBreadcrumbEntry[]>
