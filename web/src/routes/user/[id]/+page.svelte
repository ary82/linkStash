<script lang="ts">
  export let data;
</script>

<div class="pt-4 min-h-screen bg-base-200">
  <img src={data.user.picture} alt="user profile" />
  <h1>{data.user.username} @ {data.user.created_at}</h1>
  <h2>stars earned: {data.user.stars}</h2>
  <div class="divider text-xl">Public Stashes</div>
  {#if data.user.public_stashes !== null}
    {#each data.user.public_stashes as stash (stash.id)}
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
  {/if}
</div>
