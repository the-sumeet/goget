<script>
  import Sidebar from "./Sidebar.svelte";
  import Request from "./Request.svelte";
  import Response from "./Response.svelte";
  import UrlBar from "./UrlBar.svelte";
</script>

<div class="h-screen flex flex-col">
  <div class="h-full flex flex-row">
    <!-- Sidebar -->
    <Sidebar />
  
    <div class="flex flex-col w-full">
      <div class="grow-0">
        <UrlBar />
      </div>
  
      <div class="flex h-full flex-row overflow-y-auto">
        <Request />
        <Response />
      </div>
    </div>
  </div>
</div>
