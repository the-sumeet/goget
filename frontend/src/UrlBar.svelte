<script>
  import { Methods } from "../wailsjs/go/main/App";
  import { currentUrl, currentMethod, currentParams, currentResponse } from "./stores.js";
  import { SendHttpRequest } from "../wailsjs/go/main/App";
  import { onDestroy } from "svelte/internal";

  let opened = false;
  let sending = false;
  let methods = [];

  let method = "";
  let url = ""
  let params = {};

  const currentUrlSubSub = currentUrl.subscribe((value) => {
    url = value;
  });
  const currentMethodUnsub = currentMethod.subscribe((value) => {
    method = value;
  });
  const currentParamsUnsub = currentParams.subscribe((value) => {
    params = value;
  });

  $: currentUrl.set(url);



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
    SendHttpRequest(url).then((response) => {
      console.log(response);
      currentResponse.set(response);
      sending = false;
    });
  }

  

  onDestroy(() => {
    currentUrlSubSub();
    currentMethodUnsub();
    currentParamsUnsub()
  });

</script>

<div class="flex w-full items-center mb-4 p-2">
  
  <!-- Methods -->
  <div class="relative inline-block mr-2">
    <!-- Dropdown toggle button -->
    <button
      on:click={() => (opened = !opened)}
      class="btn btn-secondary relative z-10 flex items-center"
    >
      <span class="mx-1 text-md">{method.toUpperCase()}</span>
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
      bind:value={url}
      type="text"
      placeholder="https://example.com"
      class="block w-full text-md rounded-lg border border-lightDark bg-dark px-4 py-2"
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
