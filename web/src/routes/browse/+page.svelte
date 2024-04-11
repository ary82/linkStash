<script lang="ts">
  export let data;
  // async function getPublicStashes(): Promise<any> {
  //   const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/stash`, {
  //     method: "GET",
  //   });
  //
  //   if (res.ok) {
  //     const json = await res.json();
  //     return json;
  //   }
  // }
</script>

<div class="pt-4 min-h-screen bg-base-200">
  {#each data.stashes as stash (stash.id)}
    <div class="collapse collapse-arrow bg-base-200">
      <input type="radio" name="my-accordion-2" />
      <div class="collapse-title text-xl font-medium">{stash.title}</div>
      <div class="collapse-content">
        <p>{stash.body}</p>
      </div>
    </div>
    <div class="divider"></div>
  {/each}
</div>
