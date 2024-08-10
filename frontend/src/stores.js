import { writable } from 'svelte/store';

const currentUrl = writable('');
const currentResponse = writable('');

export { currentUrl, currentResponse };
