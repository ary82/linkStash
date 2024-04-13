<script lang="ts">
  import "../app.css";
  import { isAuthenticated, storedTheme } from "$lib/store";
  import { goto } from "$app/navigation";

  // Import icons
  import Github from "$lib/icons/Github.svelte";
  import Menu from "$lib/icons/Menu.svelte";
  import Mail from "$lib/icons/Mail.svelte";
  import Linkedin from "$lib/icons/Linkdein.svelte";
  import Archive from "$lib/icons/Archive.svelte";

  export let data;

  let checkboxVal: boolean = $storedTheme === "dark";
  $: {
    if (checkboxVal) {
      storedTheme.set("dark");
    } else {
      storedTheme.set("light");
    }
  }

  // Function for logging out
  // Server deletes jwt token cookie
  async function Logout(): Promise<void> {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/logout`, {
      method: "POST",
    });

    const json = await res.json();
    if (res.ok) {
      isAuthenticated.set("false");
    }
    console.log(json);
    goto("/");
  }
</script>

<div class="sticky top-0 z-[1]">
  <div class="navbar bg-base-100">
    <div class="navbar-start">
      <div class="dropdown">
        <div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
          <Menu width={20} height={20} />
        </div>
        <ul
          class="menu dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
        >
          <li><a href="/">Home</a></li>
          <li><a href="/browse">Browse</a></li>
          {#if $isAuthenticated === "true"}
            <li><a href="/stashes">Stashes</a></li>
          {/if}
        </ul>
      </div>
      <a class="btn btn-ghost text-xl" href="/"
        ><Archive width={24} height={24} />urlStash</a
      >
    </div>
    <div class="navbar-center hidden lg:flex">
      <ul class="menu menu-horizontal px-1">
        <li><a href="/">Home</a></li>
        <li><a href="/browse">Browse</a></li>
        {#if $isAuthenticated === "true"}
          <li><a href="/stashes">Stashes</a></li>
        {/if}
      </ul>
    </div>
    <div class="navbar-end mr-1">
      <input
        type="checkbox"
        value="dim"
        id="theme-controller"
        class="toggle theme-controller m-4"
        bind:checked={checkboxVal}
      />
      {#if $isAuthenticated === "true"}
        <div class="dropdown dropdown-end">
          <div
            tabindex="0"
            role="button"
            class="btn btn-ghost btn-circle avatar"
          >
            <div class="w-10 rounded-full">
              <img alt="your google profile" src={data.user.picture} />
            </div>
          </div>
          <ul
            class="menu dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
          >
            <li>
              <a href="/user/{data.user.id}" class="justify-between">
                Profile
              </a>
            </li>
            <li><a href="/">Settings</a></li>
            <li><button on:click={Logout}>Logout</button></li>
          </ul>
        </div>
      {:else}
        <a href="/login" class="btn btn-primary">Sign in</a>
      {/if}
    </div>
  </div>
</div>

<slot />

<footer class="footer p-10 bg-base-300">
  <aside>
    <h2 class="text-xl font-semibold">urlStash</h2>
    <p>Copyright © 2024 - Aryan Goyal<br />MIT License</p>
  </aside>
  <nav>
    <h6 class="footer-title">Social</h6>
    <div class="grid grid-flow-col gap-4">
      <a
        href="https://github.com/ary82/urlstash"
        target="_blank"
        rel="noopener noreferrer"><Github width={24} height={24} /></a
      >
      <a
        href="https://linkedin.com/in/aryan-goyal1"
        target="_blank"
        rel="noopener noreferrer"><Linkedin width={26} height={26} /></a
      >
      <a href="mailto:mail@ary82.dev"><Mail width={26} height={26} /></a>
    </div>
  </nav>
</footer>
