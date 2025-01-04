import { Drawer as DrawerPrimitive } from 'vaul-svelte'

import Content from './drawer-content.svelte'
import Description from './drawer-description.svelte'
import Footer from './drawer-footer.svelte'
import Header from './drawer-header.svelte'
import NestedRoot from './drawer-nested.svelte'
import Overlay from './drawer-overlay.svelte'
import ScrollArea from './drawer-scroll-area.svelte'
import Title from './drawer-title.svelte'
import Root from './drawer.svelte'
import RootByNavigationState from './root-by-navigation-state.svelte'

const Trigger: typeof DrawerPrimitive.Trigger = DrawerPrimitive.Trigger
const Portal: typeof DrawerPrimitive.Portal = DrawerPrimitive.Portal
const Close: typeof DrawerPrimitive.Close = DrawerPrimitive.Close

export {
	Close,
	Content,
	Description,
	Root as Drawer,
	RootByNavigationState as DrawerByNavigationState,
	Close as DrawerClose,
	Content as DrawerContent,
	Description as DrawerDescription,
	Footer as DrawerFooter,
	Header as DrawerHeader,
	NestedRoot as DrawerNestedRoot,
	Overlay as DrawerOverlay,
	Portal as DrawerPortal,
	ScrollArea as DrawerScrollArea,
	Title as DrawerTitle,
	Trigger as DrawerTrigger,
	Footer,
	Header,
	NestedRoot,
	Overlay,
	Portal,
	Root,
	ScrollArea,
	Title,
	Trigger,
}
