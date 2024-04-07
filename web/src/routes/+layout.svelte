<script lang="ts">
  import { isAuthenticated, storedTheme, storedPicture } from "$lib/store";
  import "../app.css";
  import { goto } from "$app/navigation";

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
  async function Logout() {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/logout`, {
      method: "POST",
      credentials: "include",
    });

    const json = await res.json();
    if (res.ok) {
      isAuthenticated.set("false");
      storedPicture.set("");
    }
    console.log(json);
    goto("/");
  }
</script>

<div class="relative z-[1]">
  <div class="fixed navbar bg-base-100">
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
          class="menu dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
        >
          <li><a href="/">Home</a></li>
          <li><a href="/browse">Browse</a></li>
          {#if $isAuthenticated === "true"}
            <li><a href="/stashes">Stashes</a></li>
          {/if}
        </ul>
      </div>
      <a class="btn btn-ghost text-xl" href="/">urlStash</a>
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
              <img alt="your google profile" src={$storedPicture} />
            </div>
          </div>
          <ul
            class="menu dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
          >
            <li>
              <a href="/" class="justify-between"> Profile </a>
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

<footer class="footer items-center p-4 bg-base-100">
  <aside class="items-center grid-flow-col">
    <svg
      width="36"
      height="36"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
      fill-rule="evenodd"
      clip-rule="evenodd"
      class="fill-current"
      ><path
        d="M22.672 15.226l-2.432.811.841 2.515c.33 1.019-.209 2.127-1.23 2.456-1.15.325-2.148-.321-2.463-1.226l-.84-2.518-5.013 1.677.84 2.517c.391 1.203-.434 2.542-1.831 2.542-.88 0-1.601-.564-1.86-1.314l-.842-2.516-2.431.809c-1.135.328-2.145-.317-2.463-1.229-.329-1.018.211-2.127 1.231-2.456l2.432-.809-1.621-4.823-2.432.808c-1.355.384-2.558-.59-2.558-1.839 0-.817.509-1.582 1.327-1.846l2.433-.809-.842-2.515c-.33-1.02.211-2.129 1.232-2.458 1.02-.329 2.13.209 2.461 1.229l.842 2.515 5.011-1.677-.839-2.517c-.403-1.238.484-2.553 1.843-2.553.819 0 1.585.509 1.85 1.326l.841 2.517 2.431-.81c1.02-.33 2.131.211 2.461 1.229.332 1.018-.21 2.126-1.23 2.456l-2.433.809 1.622 4.823 2.433-.809c1.242-.401 2.557.484 2.557 1.838 0 .819-.51 1.583-1.328 1.847m-8.992-6.428l-5.01 1.675 1.619 4.828 5.011-1.674-1.62-4.829z"
      ></path></svg
    >
    <p>Copyright © 2024 - All right reserved</p>
  </aside>
  <nav class="grid-flow-col gap-4 md:place-self-center md:justify-self-end">
    <!-- Github Button -->
    <a
      href="https://github.com/ary82/urlstash"
      target="_blank"
      rel="noopener noreferrer"
      ><svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        class="fill-current"
        ><path
          d="M 12.009929,0 C 5.3687485,0 0,5.4083277 0,12.099166 c 0,5.348345 3.4399422,9.875607 8.2120539,11.477947 0.5966369,0.120458 0.8151825,-0.260337 0.8151825,-0.580658 0,-0.280495 -0.019667,-1.241949 -0.019667,-2.243718 C 5.6666982,21.474011 4.9709906,19.310434 4.9709906,19.310434 4.4340912,17.908202 3.6385754,17.547811 3.6385754,17.547811 c -1.0934656,-0.741187 0.07965,-0.741187 0.07965,-0.741187 1.2129404,0.08014 1.8494022,1.241949 1.8494022,1.241949 1.073553,1.842764 2.8034804,1.32209 3.4994338,1.001524 0.099317,-0.781258 0.4176704,-1.322091 0.7556909,-1.622499 -2.6645847,-0.280495 -5.4680651,-1.32209 -5.4680651,-5.969564 0,-1.32209 0.4769162,-2.403756 1.232607,-3.2449968 -0.119229,-0.300408 -0.5368994,-1.5426025 0.1194749,-3.2051716 0,0 1.0140615,-0.3205663 3.3005549,1.2419487 a 11.54678,11.54678 0 0 1 3.002605,-0.4007079 c 1.014062,0 2.04779,0.1403707 3.002359,0.4007079 2.28674,-1.562515 3.300801,-1.2419487 3.300801,-1.2419487 0.656374,1.6625691 0.238458,2.9047636 0.119229,3.2051716 0.775604,0.8412408 1.232853,1.9229068 1.232853,3.2449968 0,4.647474 -2.80348,5.668911 -5.487978,5.969564 0.437583,0.38055 0.815183,1.101578 0.815183,2.243473 0,1.622498 -0.01967,2.924676 -0.01967,3.325138 0,0.320567 0.218792,0.701362 0.815183,0.58115 C 20.56,21.974527 23.999943,17.447511 23.999943,12.099166 24.019612,5.4083277 18.631197,0 12.009929,0 Z"
        ></path></svg
      ></a
    >
    <!-- Mail Button -->
    <a href="mailto:mail@ary82.dev">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        ><path
          d="M3 7a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v10a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2v-10z"
        /> <path d="M3 7l9 6l9 -6" /></svg
      ></a
    >
  </nav>
</footer>
