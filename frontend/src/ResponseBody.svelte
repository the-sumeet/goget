<script>
    let editor;

    import { currentResponse } from "./stores.js";
    import { onDestroy } from "svelte/internal";
    import { onMount } from "svelte/internal";
    import ace from "ace-builds/src-noconflict/ace"; // Import Ace from node_modules
    import "ace-builds/src-noconflict/mode-javascript"; // Import the mode you need
    import "ace-builds/src-noconflict/theme-monokai"; // Import the theme you want

    let body;

    onMount(() => {
        editor = ace.edit("editor");
        editor.setTheme("ace/theme/dracula");
        editor.session.setMode("ace/mode/javascript");
        editor.setReadOnly(true);
    });

    $: if (editor && body) {
        editor.setValue(body, -1);
    }

    const unsub = currentResponse.subscribe((value) => {
        body = value.body;
    });

    onDestroy(unsub);
</script>

<div class="h-full text-base" id="editor"></div>
