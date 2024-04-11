import type { PageLoad } from "./$types";

type stash = {
  id: number;
  author: string;
  title: string;
  body: string;
  stars: number;
  created_at: string;
};

type s = {
  stashes: stash[];
};

export const load: PageLoad = async ({ fetch }: any): Promise<any> => {
  try {
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/stash`);
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const stashes = (await response.json()) as s;
    console.log(stashes)
    return { stashes };
  } catch (error) {
    console.error(error);
    throw new Error(`Unable to fetch`);
  }
};
