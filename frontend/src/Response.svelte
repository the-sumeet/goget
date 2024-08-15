<script>
  import { each, onDestroy } from "svelte/internal";
  import { currentResponse } from "./stores.js";
  import { onMount } from "svelte/internal";
  import ace from "ace-builds/src-noconflict/ace"; // Import Ace from node_modules
  import "ace-builds/src-noconflict/mode-javascript"; // Import the mode you need
  import "ace-builds/src-noconflict/theme-monokai"; // Import the theme you want

  let response;
  let editor;

  const unsub = currentResponse.subscribe((value) => {
    response = value;
  });

  $: if (editor) {
    editor.setValue(response);
  }

  onMount(() => {
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/dracula");
    editor.session.setMode("ace/mode/javascript");
    editor.setValue(response);
  });

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
      <button class="btn btn-text">
        <i class="bi bi-copy"></i>
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
  <div class="flex p-2">
    <!-- Copy button -->
    <button class="btn btn-secondary">
      <i class="bi bi-copy"></i>
    </button>

    <!-- Format button -->
    <button class="btn btn-secondary">
      <i class="bi bi-magic"></i>
    </button>
  </div>
</div>
