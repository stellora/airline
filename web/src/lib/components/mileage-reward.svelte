<script lang="ts">
  import { Badge } from './ui/badge';
  import Plane from 'lucide-svelte/icons/plane';

  export let miles: number;
</script>

<div class="flex items-center gap-1" title="Mileage reward for this flight">
  <Plane class="h-3 w-3" />
  <Badge variant="outline" class="text-xs py-0">
    {miles.toLocaleString()} miles
  </Badge>
</div>