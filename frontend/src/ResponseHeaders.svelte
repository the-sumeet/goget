<script>
    let editor;

    import { currentResponse } from "./stores.js";
    import { each, onDestroy } from "svelte/internal";
    let header;

    $: if (editor && header) {
        editor.setValue(JSON.stringify(header, null, 2), -1);
    }

    const unsub = currentResponse.subscribe((value) => {
        header = value.header;
    });

    onDestroy(unsub);
</script>

{#if header}
<div class="h-fit">
    <dl class="-my-3 text-sm">
        {#each Object.entries(header) as [name, values]}
        {#each values as value}
        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">{name}</dt>
            <dd class="text-gray-700 sm:col-span-2">{value}</dd>
        </div>
        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">{name}</dt>
            <dd class="text-gray-700 sm:col-span-2">{value}</dd>
        </div>
        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">{name}</dt>
            <dd class="text-gray-700 sm:col-span-2">{value}</dd>
        </div>
        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">{name}</dt>
            <dd class="text-gray-700 sm:col-span-2">{value}</dd>
        </div>
        {/each}
        {/each}
    </dl>
</div>
{/if}