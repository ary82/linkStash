import { browser } from "$app/environment";
import { writable } from "svelte/store";


export const isAuthenticated = writable(browser && localStorage.getItem("auth") || "false")
isAuthenticated.subscribe((val) => {
  if (browser) return (localStorage.setItem("auth", val === "true" ? "true" : "false"))
})

export const storedTheme = writable(browser && localStorage.getItem("theme") || "false")
storedTheme.subscribe((val) => {
  if (browser) return (localStorage.setItem("theme", val === "dark" ? "dark" : "light"))
})

export const storedPicture = writable(browser && localStorage.getItem("picture") || "")
storedPicture.subscribe((val) => {
  if (browser) return (localStorage.setItem("picture", val))
})
