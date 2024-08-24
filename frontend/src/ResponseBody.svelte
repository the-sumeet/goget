<script>
    let editor;

    import { currentResponse } from "./stores.js";
    import { onDestroy } from "svelte/internal";
    import { onMount } from "svelte/internal";
    import * as ace from 'brace';
    import 'brace/mode/html';
    import 'brace/mode/json';
    import 'brace/mode/text';
    import 'brace/mode/xml';
    import 'brace/mode/yaml';
    import 'brace/theme/twilight';

    import { currentResponseType } from "./stores.js";

    let body;
    let responseType = "text";

    onMount(() => {
        editor = ace.edit("editor");
        editor.setTheme("ace/theme/twilight");
        editor.setReadOnly(true);
        editor.session.setMode("ace/mode/" + responseType);
        editor.setFontSize('14px');

    });

    $: if (editor && body) {
        editor.setValue(body, -1);
    }

    // $: if (responseType) {editor.session.setMode("ace/mode/"+responseType)};

    const unsub = currentResponse.subscribe((value) => {
        body = value.body;
    });

    const unsubResponseType = currentResponseType.subscribe((value) => {
        responseType = value;
        if (editor) {
            console.log("Setting mode to " + responseType);
            editor.session.setMode("ace/mode/" + responseType)
        }
    });

    onDestroy(() => {
        unsub();
        unsubResponseType();
    });
</script>

<div class="h-full text-base" id="editor"></div>
