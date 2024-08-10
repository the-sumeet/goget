<script>
  import Sidebar from "./Sidebar.svelte";
  import Request from "./Request.svelte";
  import RequestOut from "./Response.svelte";
  import TopBar from "./UrlBar.svelte";
  // import { methods } from "./constants";
  import { Methods } from "../wailsjs/go/main/App";
  import { currentUrl, currentResponse } from "./stores.js";
  import { SendHttpRequest } from "../wailsjs/go/main/App";

  let opened = false;
  let sending = false;
  let methods = [];
  let method = "";

  Methods().then((response) => {
    methods = response;
    method = methods[0];
  });

  function setMethod(newMethod) {
    method = newMethod;
    opened = !opened;
  }

  function sendRequest() {
    if (sending) {
      return;
    }
    sending = true;
    SendHttpRequest($currentUrl).then((response) => {
      currentResponse.set(response);
      sending = false;
    });
  }
</script>

<div class="flex w-full items-center mb-4 p-2">
  
  <!-- Methods -->
  <div class="relative inline-block mr-2">
    <!-- Dropdown toggle button -->
    <button
      on:click={() => (opened = !opened)}
      class="btn btn-secondary relative z-10 flex items-center"
    >
      <span class="mx-1 text-xl">{method}</span>
    </button>

    <!-- Dropdown menu -->
    {#if opened}
      <div
        class="absolute right-0 z-20 py-2 mt-2 overflow-hidden origin-top-right bg-white rounded-md shadow-xl bg-dark border border-lightDark"
      >
        {#each methods as method, i}
          <a
            on:click={() => setMethod(method)}
            href="#!"
            class="block px-4 py-3 hover:bg-lightDark"
          >
            {method}
          </a>
        {/each}
      </div>
    {/if}
  </div>

  <!-- Input -->
  <div class="w-full mr-2">
    <input
      bind:value={$currentUrl}
      type="text"
      placeholder="https://example.com"
      class="block w-full text-xl rounded-lg border border-lightDark bg-dark px-4 py-2"
    />
  </div>

  <!-- Send Button -->
  {#if sending}
    <button
      class=" min-w-20 btn btn-primary cursor-no-drop hover:bg-accent"
      on:click={sendRequest}
    >
      <div class="text-xl animate-spin"><i class="bi bi-arrow-repeat"></i></div>
    </button>
  {:else}
    <button class="min-w-20 text-xl btn btn-primary" on:click={sendRequest}>
      <i class="bi bi-send-fill"></i>
    </button>
  {/if}
</div>
