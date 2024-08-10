<script>
  import Sidebar from "./Sidebar.svelte";
  import Request from "./Request.svelte";
  import Response from "./Response.svelte";
  import UrlBar from "./UrlBar.svelte";
  import { currentUrl, currentResponse } from "./stores.js";
  import { SendHttpRequest } from "../wailsjs/go/main/App";
</script>

<div class="flex h-screen">
  <!-- Sidebar -->
  <Sidebar />

  <div class="flex flex-col w-full">
    <!-- URL Bar -->
    <UrlBar />

    <div class="flex grow">
      <Request />
      <Response />
    </div>
  </div>
</div>
