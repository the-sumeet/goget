<script>
  import { each, onDestroy } from "svelte/internal";
  import { currentResponse } from "./stores.js";
  import { onMount } from "svelte/internal";
  import ace from "ace-builds/src-noconflict/ace"; // Import Ace from node_modules
  import "ace-builds/src-noconflict/mode-javascript"; // Import the mode you need
  import "ace-builds/src-noconflict/theme-monokai"; // Import the theme you want
  import { formatDuration } from "./utils.js";
  import {ClipboardSetText} from '../wailsjs/runtime/runtime.js';
  
  let status;
  let body;
  let time;
  let editor;
  let copying = false;

  onMount(() => {
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/dracula");
    editor.session.setMode("ace/mode/javascript");
    editor.setValue("");
  });

  const unsub = currentResponse.subscribe((value) => {
    console.log(value);
    status = value.status;
    body = value.body;
    time = value.time;
    console.log(body);
  });

  $: if (editor && body) {
    editor.setValue(body);
  }



  function onCopy() {
    if (copying) {
      return;
    }
    copying = true;
    ClipboardSetText(editor.getValue()).then((copied) => {
      copying = false;
    });
  }

  onDestroy(unsub);
</script>

<div class="basis-1/2 flex flex-col">
  <!-- Response Tabs -->
  <div
    class="flex justify-between overflow-x-auto overflow-y-hidden border-b border-lightDark whitespace-nowrap"
  >
    <!-- Tab Names -->
    <div>
      <button
        class="inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-blue-600 bg-transparent border-b-4 border-accent sm:text-base dark:border-blue-400 dark:text-blue-300 whitespace-nowrap focus:outline-none"
      >
        Response
      </button>

      <button
        class="inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-gray-700 bg-transparent border-b-2 border-transparent sm:text-base dark:text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400"
      >
        Headers
      </button>
    </div>

    <!-- Right Side Buttons -->
    <div class="flex">
      <!-- Copy button -->
      <button on:click={onCopy} class="btn btn-text">
        {#if copying}
        <div class="animate-spin"><i class="bi bi-arrow-repeat"></i></div>
        {:else}
        <i class="bi bi-copy"></i>
        {/if}
      </button>

      <!-- Format button -->
      <button class="btn btn-text">
        <i class="bi bi-magic"></i>
      </button>
      
    </div>
  </div>

  <!-- Outout Text -->
  <div class="grow text-base w-full" id="editor">sidfniefni</div>

  <!-- Bottom Toolbar -->
  <div class="flex p-2 space-x-2">
    {#if status && status.startsWith("2")}
      <span class="text-green">{status}</span>
    {:else if status && status.startsWith("4")}
      <span class="text-yellow">{status}</span>
    {:else if status}
      <span class="text-red">{status}</span>
    {/if}

    {#if time}
      <span class="text-gray-500">{formatDuration(time)}</span>
    {/if}
  </div>
</div>
