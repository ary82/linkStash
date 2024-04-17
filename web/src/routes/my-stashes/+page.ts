import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }: any) => {
  try {
    const response = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/my-stashes`,
      {
        credentials: "include",
      },
    );
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const stashes = (await response.json()) as stash[];
    return { stashes };
  } catch (error) {
    console.error(error);
    throw new Error(`Unable to fetch`);
  }
};
