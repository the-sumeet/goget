<script>
    let editor;

    import { currentResponse } from "./stores.js";
    import { each, onDestroy } from "svelte/internal";
    import { ClipboardSetText } from "../wailsjs/runtime/runtime.js";
    import { onMount } from "svelte/internal";
    import ace from "ace-builds/src-noconflict/ace"; // Import Ace from node_modules
    import "ace-builds/src-noconflict/mode-javascript"; // Import the mode you need
    import "ace-builds/src-noconflict/theme-monokai"; // Import the theme you want
    import { edit } from "ace-builds";

    let header;

    onMount(() => {
        editor = ace.edit("editor");
        editor.setTheme("ace/theme/dracula");
        editor.session.setMode("ace/mode/javascript");
        editor.setValue("");
        editor.setReadOnly(true);
    });

    $: if (editor && header) {
        editor.setValue(JSON.stringify(header, null, 2), -1);
    }

    const unsub = currentResponse.subscribe((value) => {
        header = value.header;
    });

    onDestroy(unsub);
</script>

<div class="flow-root rounded-lg py-3">
    <dl class="-my-3 divide-y divide-gray-100 text-sm">
        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">Title</dt>
            <dd class="text-gray-700 sm:col-span-2">Mr</dd>
        </div>

        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">Name</dt>
            <dd class="text-gray-700 sm:col-span-2">John Frusciante</dd>
        </div>

        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">Occupation</dt>
            <dd class="text-gray-700 sm:col-span-2">Guitarist</dd>
        </div>

        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">Salary</dt>
            <dd class="text-gray-700 sm:col-span-2">$1,000,000+</dd>
        </div>

        <div
            class="grid grid-cols-1 gap-1 p-3 even:bg-lightDark sm:grid-cols-3 sm:gap-4"
        >
            <dt class="font-medium text-gray-900">Bio</dt>
            <dd class="text-gray-700 sm:col-span-2">
                Lorem ipsum dolor, sit amet consectetur adipisicing elit. Et
                facilis debitis explicabo doloremque impedit nesciunt dolorem
                facere, dolor quasi veritatis quia fugit aperiam aspernatur
                neque molestiae labore aliquam soluta architecto?
            </dd>
        </div>
    </dl>
</div>
