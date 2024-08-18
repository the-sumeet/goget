<script>
  import { onDestroy } from "svelte/internal";
  import { currentResponse } from "./stores.js";
  import "ace-builds/src-noconflict/mode-javascript"; // Import the mode you need
  import "ace-builds/src-noconflict/theme-monokai"; // Import the theme you want
  import { formatDuration } from "./utils.js";
  import { ClipboardSetText } from "../wailsjs/runtime/runtime.js";
  import ResponseBody from "./ResponseBody.svelte";
  import ResponseHeaders from "./ResponseHeaders.svelte";

  let status;
  let time;
  let header;
  let editor;
  let copying = false;
  let currentTab = 0;
  let selectedTabCss =
    "inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-blue-600 bg-transparent border-b-4 border-accent sm:text-base dark:border-blue-400 dark:text-blue-300 whitespace-nowrap focus:outline-none";
  let unSelectedTabCss =
    "inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-gray-700 bg-transparent border-b-2 border-transparent sm:text-base dark:text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400";

  const unsub = currentResponse.subscribe((value) => {
    status = value.status;
    header = value.header;
    time = value.time;
  });

  function onCopy() {
    if (copying) {
      return;
    }
    copying = true;
    ClipboardSetText(editor.getValue()).then((copied) => {
      copying = false;
    });
  }

  function setTab(index) {
    currentTab = index;
  }

  onDestroy(unsub);
</script>

<div class="basis-1/2 flex flex-col overflow-y-auto">
  <!-- Response Tabs -->
  <div
    class="flex justify-between overflow-x-auto overflow-y-hidden border-b border-lightDark whitespace-nowrap"
  >
    <!-- Tab Names -->
    <div>
      <button
        on:click={() => setTab(0)}
        class={currentTab == 0 ? selectedTabCss : unSelectedTabCss}
      >
        Response
      </button>

      <button
        on:click={() => setTab(1)}
        class={currentTab == 1 ? selectedTabCss : unSelectedTabCss}
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

  <!-- Response body or headers -->
  <div class="flex-1">
    {#if currentTab === 0}
      <ResponseBody />
    {:else}
      <ResponseHeaders />
    {/if}
  </div>

  <!-- Bottom Toolbar -->
  {#if status || time}
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
  {/if}
</div>
