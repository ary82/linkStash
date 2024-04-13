import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }: any) => {
  try {
    const response = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/user/${params.id}`,
    );
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const user = (await response.json()) as userDetail;
    return { user };
  } catch (error) {
    console.error(error);
    throw new Error(`Unable to fetch`);
  }
};
