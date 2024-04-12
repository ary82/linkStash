<script lang="ts">
  export let data;
</script>

<div class="pt-4 min-h-screen bg-base-200">
  <a href="/user/{data.stash.author_id}">MORE FROM USER</a>
  <h1>{data.stash.author}: {data.stash.title} @ {data.stash.created_at}</h1>
  <p>{data.stash.body}</p>
  <p>Stars: {data.stash.stars}</p>
  <p>Is Public? {data.stash.is_public}</p>
  <div class="divider">Links</div>
  {#if data.stash.links !== null}
    {#each data.stash.links as link (link.id)}
      <h2>{link.url}</h2>
      <h2>{link.comment}</h2>
    {/each}
  {/if}
  <div class="divider">Comments</div>
  {#if data.stash.comments !== null}
    {#each data.stash.comments as comment (comment.id)}
      <p>{comment.author}: {comment.body} @ {comment.created_at}</p>
      <a href="/user/{comment.author_id}">SEE USER PROFILE</a>
    {/each}
  {/if}
</div>
