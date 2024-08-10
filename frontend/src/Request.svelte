<script>
  import "bootstrap-icons/font/bootstrap-icons.css";

  import { currentUrl, currentResponse } from "./stores.js";
  import { SendHttpRequest } from "../wailsjs/go/main/App";

  let sending = false;
  let opened = false;


  function onUrlUpdate(event) {
    currentUrl.set(event.target.value);
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

<div class="basis-1/2 flex flex-col border-r border-lightDark">
  <!-- Tabs -->
  <div
    class="flex  overflow-x-auto overflow-y-hidden border-b border-lightDark whitespace-nowrap"
  >
    <button
      class="inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-blue-600 bg-transparent border-b-4 border-accent sm:text-base dark:border-blue-400 dark:text-blue-300 whitespace-nowrap focus:outline-none"
    >
      Params
    </button>

    <button
      class="inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-gray-700 bg-transparent border-b-2 border-transparent sm:text-base dark:text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400"
    >
      Headers
    </button>

    <button
      class="inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-gray-700 bg-transparent border-b-2 border-transparent sm:text-base dark:text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400"
    >
      Auth
    </button>
    <button
      class="inline-flex items-center h-10 px-4 -mb-px text-sm text-center text-gray-700 bg-transparent border-b-2 border-transparent sm:text-base dark:text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400"
    >
      Body
    </button>
  </div>
</div>
