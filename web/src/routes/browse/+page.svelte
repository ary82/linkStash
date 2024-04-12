<script lang="ts">
  export let data;
</script>

<div class="pt-4 min-h-screen bg-base-200">
  {#each data.stashes as stash (stash.id)}
    <a href="/user/{stash.author_id}">MORE FROM USER</a>
    <div class="collapse collapse-arrow bg-base-200">
      <input type="radio" name="my-accordion-2" />
      <div class="collapse-title text-xl font-medium">
        {stash.author}: {stash.title} @ {stash.created_at}
        <div>Stars: {stash.stars}</div>
      </div>
      <div class="collapse-content">
        <p>{stash.body}</p>
      </div>
    </div>
    <a href="/stash/{stash.id}">SEE DETAILS</a>
    <div class="divider"></div>
  {/each}
</div>
