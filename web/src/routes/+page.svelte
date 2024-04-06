<script lang="ts">
  import { onMount } from "svelte";

  async function sendJWT(googleRespone: any) {
    console.log(googleRespone.credential);
    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/login`, {
      method: "POST",
      body: googleRespone.credential,
    });

    const json = await res.json();
    const result = JSON.stringify(json);
    console.log(result);
  }

  onMount(() => {
    console.log("the component has mounted");
    google.accounts.id.initialize({
      client_id:
        "274396243883-06ihk2heb22490spe6hdtc0115a5jn57.apps.googleusercontent.com",
      callback: sendJWT,
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
    // google.accounts.id.prompt();
  });
</script>

<div class="navbar bg-base-100">
  <div class="navbar-start">
    <div class="dropdown">
      <div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          ><path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h8m-8 6h16"
          /></svg
        >
      </div>
      <ul
        tabindex="0"
        class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
      >
        <li><a>Browse</a></li>
        <li>
          <a>Parent</a>
          <ul class="p-2">
            <li><a>Submenu 1</a></li>
            <li><a>Submenu 2</a></li>
          </ul>
        </li>
        <li><a>Item 3</a></li>
      </ul>
    </div>
    <a class="btn btn-ghost text-xl">urlStash</a>
  </div>
  <div class="navbar-center hidden lg:flex">
    <ul class="menu menu-horizontal px-1">
      <li><a>Browse</a></li>
      <li>
        <details>
          <summary>Parent</summary>
          <ul class="p-2">
            <li><a>Submenu 1</a></li>
            <li><a>Submenu 2</a></li>
          </ul>
        </details>
      </li>
      <li><a>Item 3</a></li>
    </ul>
  </div>
  <div class="navbar-end mr-1">
    <input
      type="checkbox"
      value="dim"
      id="theme-controller"
      class="toggle theme-controller m-4"
    />
    <div id="googleSigninButton" />
  </div>
</div>

<div class="hero min-h-screen bg-base-200">
  <div class="hero-content text-center">
    <div class="max-w-md">
      <h1 class="text-5xl font-bold">Hello there</h1>
      <p class="py-6">
        Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
        excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a
        id nisi.
      </p>
      <button class="btn btn-primary">Get Started</button>
    </div>
  </div>
</div>
