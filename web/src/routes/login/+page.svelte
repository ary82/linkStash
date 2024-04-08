<script lang="ts">
  import { goto } from "$app/navigation";
  import { isAuthenticated, storedPicture } from "$lib/store";
  import { onMount } from "svelte";

  // Function for logging in
  // Called on pressing login with Google button
  // Server sets a jwt token as cookie
  async function Login(googleRespone: any) {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/login`, {
      method: "POST",
      credentials: "include",
      body: googleRespone.credential,
    });

    const json = await res.json();
    if (res.ok) {
      isAuthenticated.set("true");
      storedPicture.set(json.picture);
      goto("/");
    }
  }

  onMount(() => {
    if ($isAuthenticated === "false") {
      google.accounts.id.initialize({
        client_id:
          "274396243883-06ihk2heb22490spe6hdtc0115a5jn57.apps.googleusercontent.com",
        callback: Login,
        prompt_parent_id: "parent_id",
        use_fedcm_for_prompt: true,
      });
      google.accounts.id.renderButton(
        document.getElementById("googleSigninButton"),
        {
          size: "large",
          width: "100",
        },
      );
    }
  });
</script>

<div class="hero min-h-screen bg-base-200">
  <div class="hero-content text-center">
    <div class="max-w-md">
      <h1 class="text-5xl font-bold">Get Started</h1>
      <div class="card bg-base-100 shadow-xl mt-10">
        <div class="card-body items-center text-center">
          <h2 class="text-2xl card-title">urlStash</h2>
          <p class="m-4">Sign in or Sign up now with Google</p>
          <div class="card-actions">
            <button id="googleSigninButton" />
            {#if $isAuthenticated === "true"}
              <a href="/" class="btn btn-primary">You're already logged in</a>
            {/if}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
