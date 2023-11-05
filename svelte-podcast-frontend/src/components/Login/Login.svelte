<!-- src/components/Login/Login.svelte -->
<script>
  import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from "firebase/auth";
  import { auth } from '../../firebase.js'; // Adjust the relative path as necessary
  import { user } from '../../stores/userStore';

  let isSignUp = false; // Flag to toggle between login and sign up
  let email = '';
  let password = '';
  let confirmPassword = ''; // Only used for sign up to confirm password
  let errorMessage = ''; // To display any error message from Firebase

  function toggleForm() {
    isSignUp = !isSignUp;
    errorMessage = ''; // Clear error message when toggling
  }

  async function handleLogin() {
    try {
      const response = await signInWithEmailAndPassword(auth, email, password);
      user.set(response.user);
    } catch (error) {
      console.error('Login failed:', error);
      errorMessage = error.message;
    }
  }

  async function handleSignUp() {
    if (password !== confirmPassword) {
      errorMessage = 'Passwords do not match!';
      return;
    }
    try {
      const response = await createUserWithEmailAndPassword(auth, email, password);
      user.set(response.user);
    } catch (error) {
      console.error('Sign up failed:', error);
      errorMessage = error.message;
    }
  }
</script>
  
  <div>
    <button on:click={toggleForm}>
      {#if isSignUp}
        Already have an account? Sign In
      {:else}
        Need an account? Sign Up
      {/if}
    </button>
    
    <!-- Display error message if there is one -->
    {#if errorMessage}
      <p class="error">{errorMessage}</p>
    {/if}
  
    {#if isSignUp}
      <h2>Sign Up</h2>
      <input type="email" bind:value={email} placeholder="Email" />
      <input type="password" bind:value={password} placeholder="Password" />
      <input type="password" bind:value={confirmPassword} placeholder="Confirm Password" />
      <button on:click={handleSignUp}>Sign Up</button>
    {:else}
      <h2>Sign In</h2>
      <input type="email" bind:value={email} placeholder="Email" />
      <input type="password" bind:value={password} placeholder="Password" />
      <button on:click={handleLogin}>Sign In</button>
    {/if}
  </div>
  
  <style>
    .error {
      color: red;
    }
    /* Add more CSS if needed */
  </style>
  