<script lang="ts">
  import { Badge } from '$lib/components/ui/badge';
  import AirlineCode from '$lib/components/airline-code.svelte';
  import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '$lib/components/ui/card';
  import type { AirlineLoyalty } from '$lib/airline.openapi';

  export let loyalty: AirlineLoyalty;
</script>

<Card class="hover:shadow-md transition-shadow duration-200">
  <CardHeader>
    <div class="flex justify-between items-center mb-1">
      <CardTitle class="flex items-center">
        <AirlineCode code={loyalty.airline.iataCode} /> {loyalty.airline.name}
      </CardTitle>
      <Badge variant="outline">{loyalty.mileageBalance.toLocaleString()} miles</Badge>
    </div>
    <CardDescription>Member ID: {loyalty.passengerId}</CardDescription>
  </CardHeader>
  <CardContent>
    <div class="text-sm text-muted-foreground">
      <div class="mt-2 flex justify-between">
        <span>Status:</span>
        <Badge variant="secondary">
          {#if loyalty.mileageBalance >= 100000}
            Platinum
          {:else if loyalty.mileageBalance >= 50000}
            Gold
          {:else if loyalty.mileageBalance >= 25000}
            Silver
          {:else}
            Basic
          {/if}
        </Badge>
      </div>
    </div>
  </CardContent>
</Card>