import type { PageServerLoad } from './$types';
import { fetchFromApi } from '$lib/api';
import type { Passenger } from '$lib/airline.openapi';

export const load: PageServerLoad = async ({ fetch }) => {
  // Get all passengers to populate the dropdown
  const passengers = await fetchFromApi<Passenger[]>(fetch, '/passengers');

  return {
    passengers,
    title: 'Loyalty Programs',
    breadcrumbs: [
      {
        title: 'Home',
        href: '/',
      },
      {
        title: 'Loyalty Programs',
        href: '/loyalty',
      },
    ],
  };
};