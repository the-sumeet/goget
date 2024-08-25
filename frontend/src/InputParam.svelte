<script>
    export let name;
    export let value;

</script>
<div class="flex gap-2 items-center justify-center">
    <input bind:value={name} type="text" placeholder="John Doe" class="text-sm block mt-2 w-full rounded-md border-2 border-lightDark bg-dark px-2 py-2" />
    <input bind:value={value} type="text" placeholder="John Doe" class="text-sm block mt-2 w-full rounded-md border-2 border-lightDark bg-dark px-2 py-2" />
    <input type="checkbox" name="" id="">
    <button class=" btn-text px-0"><i class="bi bi-trash"></i></button>
 </div>