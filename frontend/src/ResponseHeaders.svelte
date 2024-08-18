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

<div class="grow text-base w-full" id="editor"></div>
