<script>
	import { onMount, onDestroy } from 'svelte'; // Import onDestroy here as well
	import { user } from './stores/userStore'; // Import your user store
	import Login from './components/Login/Login.svelte';
	import Player from './components/Player/Player.svelte';
	
	console.log("hello"); // Moved this line down after imports

	let isLoggedIn = false;
  
	// Subscribe to user store
	const unsubscribe = user.subscribe(value => {
		isLoggedIn = value !== null; // Assuming null means not logged in
	});
  
	onMount(() => {
		// If you need to do any checks or setups on mount
		// You can also perform a check here if the user is already logged in, for example by checking local storage
	});
  
	// Cleanup if App is destroyed
	onDestroy(() => {
		unsubscribe();
	});
</script>



{#if isLoggedIn}
	<Player />
{:else}
	<Login />
{/if}
