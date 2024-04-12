import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }: any) => {
  try {
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/stash`);
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const stashes = (await response.json()) as stash[];
    console.log(stashes)
    return { stashes };
  } catch (error) {
    console.error(error);
    throw new Error(`Unable to fetch`);
  }
};
