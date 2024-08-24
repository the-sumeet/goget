<script>
    import InputParam from "./InputParam.svelte";
    import { onDestroy } from "svelte";
    import { currentParams } from "./stores.js";

    let params = [];

    const currentParamsUnsub = currentParams.subscribe((value) => {
        params = value;
    });

    onDestroy(() => {
        currentParamsUnsub();
    });
</script>
<div class="flex flex-col">
    <div class="flex-1 overflow-y-auto">
        <!-- Single Input -->
        {#each params as param, index}
        <InputParam />
        {/each}
    </div>
     

</div>