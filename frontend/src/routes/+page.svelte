<script>
    import {onMount} from "svelte"
    import {writable, derived} from "svelte/store"
    import axios from "axios"

    export const apiData = writable([])

    export const logIds = derived(apiData, ($apiData) => {
        console.log($apiData);
        if (!$apiData) {
            return 0;
        }

        if ($apiData.id) {
            return $apiData.id;
        }

        return 0;
    })

    console.log("hi");
    axios.post("http://localhost:8080/all")
    .then(data => {
        console.log("hello")
        console.log(data.data);
        apiData.set(data.data);
    })
    .catch(err => {
        console.log(err);
        return []
    });

    let a = 1

    function handleClick() {
        a += 1;
    }
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
{#each $logIds as logId}
    <h1>{logId}</h1>
{/each}
<p>{a}</p>
<button on:click={handleClick}>click me</button>
