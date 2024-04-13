import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }: any) => {
  try {
    const response = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/stash/${params.id}`,
    );
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const stash = (await response.json()) as stashDetail;
    return { stash };
  } catch (error) {
    console.error(error);
    throw new Error(`Unable to fetch`);
  }
};
