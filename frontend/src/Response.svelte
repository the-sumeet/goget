<script>
  import { onDestroy } from "svelte/internal";
  import { currentResponse } from "./stores.js";
  import { formatDuration } from "./utils.js";
  import ResponseBody from "./ResponseBody.svelte";
  import ResponseHeaders from "./ResponseHeaders.svelte";
  import { copyText } from "./utils.js";
  import { responseTypes } from "./constants.js";
  import { currentResponseType } from "./stores.js";

  let status;
  let time;
  let body;
  let copying = false;
  let currentTab = 0;
  let selectedTabCss =
    "inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-blue-600 bg-transparent border-b-4 border-accent sm:text-base dark:border-blue-400 dark:text-blue-300 whitespace-nowrap focus:outline-none";
  let unSelectedTabCss =
    "inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-gray-700 bg-transparent border-b-2 border-transparent sm:text-base dark:text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400";
  let responseType;
  let responseTypeDropdownOpened = false;

  const unsub = currentResponse.subscribe((value) => {
    status = value.status;
    time = value.time;
    body = value.body;
  });

  const unsubResponseType = currentResponseType.subscribe((value) => {
    responseType = value;
  });

  function setTab(index) {
    currentTab = index;
  }

  function copy() {
    copying = true;
    copyText(body).then(() => {
      copying = false;
    });
  }

  function toggleDropdown() {
    responseTypeDropdownOpened = !responseTypeDropdownOpened;
  }

  function setResponseType(type) {
    currentResponseType.set(type);
    toggleDropdown();
  }

  onDestroy(() => {
    unsub();
    unsubResponseType();
  });
</script>

<div class="basis-1/2 flex flex-col flex-none overflow-y-auto">
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
      <button on:click={copy} class="btn btn-text">
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
  <div class="flex-1 overflow-y-auto">
    {#if currentTab === 0}
      <ResponseBody />
    {:else}
      <ResponseHeaders />
    {/if}
  </div>

  <!-- Bottom Toolbar -->
  {#if status || time}
    <div class="flex p-2 space-x-4">
      <!-- Response Status -->
      {#if status && status.startsWith("2")}
        <span class="text-green">{status}</span>
      {:else if status && status.startsWith("4")}
        <span class="text-yellow">{status}</span>
      {:else if status}
        <span class="text-red">{status}</span>
      {/if}

      <!-- Response Time -->
      {#if time}
        <span class="text-gray-500">{formatDuration(time)}</span>
      {/if}

      <!-- Response Type -->
      <div class="relative inline-block">
        <button
          on:click={toggleDropdown}
          class="relative z-10 block text-gray-700 bg-white border border-transparent rounded-md dark:text-white focus:border-blue-500 focus:ring-opacity-40 dark:focus:ring-opacity-40 focus:ring-blue-300 dark:focus:ring-blue-400 focus:ring dark:bg-dark focus:outline-none"
        >
          {responseType.toUpperCase()} <i class="bi bi-chevron-up"></i>
        </button>

        {#if responseTypeDropdownOpened}
          <div
            class="absolute bottom-0 z-20 py-2 mt-2 origin-top-right bg-white rounded-md shadow-xl border border-lightDark bg-dark"
          >
            {#each responseTypes as type}
              <button
                on:click={() => setResponseType(type)}
                class="hover:bg-lightDark w-full block px-4 py-3 text-sm text-gray-600 capitalize transition-colors duration-300 transform dark:text-gray-300 bg-dark hover:bg-gray-700 hover:text-white"
                >{type.toUpperCase()}</button
              >
            {/each}
          </div>
        {/if}
      </div>
    </div>
  {/if}
</div>
