<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { Badge } from '$lib/components/ui/badge';
  import LoyaltyProgramCard from '$lib/components/loyalty-program-card.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import CircleUser from 'lucide-svelte/icons/circle-user';
  import type { AirlineLoyalty, Passenger } from '$lib/airline.openapi';

  let loyalties: AirlineLoyalty[] = [];
  let passengers: Passenger[] = [];
  let selectedPassengerId: number | null = null;
  let loading = true;
  let error = '';

  async function fetchPassengers() {
    try {
      const response = await fetch('/api/passengers');
      if (!response.ok) throw new Error('Failed to fetch passengers');
      passengers = await response.json();
    } catch (err) {
      console.error('Error fetching passengers:', err);
      error = err.message;
    }
  }

  async function fetchLoyaltyPrograms(passengerId: number) {
    loading = true;
    error = '';
    try {
      const response = await fetch(`/api/passengers/${passengerId}/loyalty-programs`);
      if (!response.ok) {
        if (response.status === 404) {
          loyalties = [];
        } else {
          throw new Error('Failed to fetch loyalty programs');
        }
      } else {
        loyalties = await response.json();
      }
    } catch (err) {
      console.error('Error fetching loyalty programs:', err);
      error = err.message;
    } finally {
      loading = false;
    }
  }

  function handlePassengerSelect(event: Event) {
    const select = event.target as HTMLSelectElement;
    const passengerId = parseInt(select.value, 10);
    if (!isNaN(passengerId)) {
      selectedPassengerId = passengerId;
      fetchLoyaltyPrograms(passengerId);
    }
  }

  onMount(async () => {
    await fetchPassengers();
    loading = false;
  });
</script>

<svelte:head>
  <title>Loyalty Programs</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex items-center justify-between">
    <h1 class="text-2xl font-bold tracking-tight">
      <div class="flex items-center gap-2">
        <CircleUser class="h-6 w-6" />
        Loyalty Programs
      </div>
    </h1>
  </div>

  <div class="rounded-md border p-4">
    <div class="mb-4">
      <Label for="passengerSelect">Select Passenger</Label>
      <select
        id="passengerSelect"
        class="w-full p-2 border rounded-md"
        on:change={handlePassengerSelect}
      >
        <option value="">Select a passenger</option>
        {#each passengers as passenger}
          <option value={passenger.id}>{passenger.name}</option>
        {/each}
      </select>
    </div>

    {#if loading}
      <div class="flex justify-center p-6">
        <div class="animate-spin rounded-full h-10 w-10 border-t-2 border-b-2 border-gray-900"></div>
      </div>
    {:else if error}
      <div class="text-red-500 p-4 rounded-md border border-red-200 bg-red-50">
        {error}
      </div>
    {:else if selectedPassengerId !== null}
      <div class="space-y-4">
        <h2 class="text-xl font-semibold">Your Loyalty Programs</h2>
        {#if loyalties.length === 0}
          <div class="p-4 text-center text-muted-foreground border rounded-md">
            No loyalty programs found for this passenger.
          </div>
        {:else}
          <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            {#each loyalties as loyalty}
              <LoyaltyProgramCard {loyalty} />
            {/each}
          </div>
        {/if}
      </div>
    {:else}
      <div class="p-4 text-center text-muted-foreground border rounded-md">
        Please select a passenger to view loyalty programs.
      </div>
    {/if}
  </div>
</div>