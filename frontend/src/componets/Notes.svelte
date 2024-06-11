<script lang="ts">
	import { handleMetaEnterKeydown } from '../utils/keydown.util';
	import { GetNotes, AddNote } from '../../wailsjs/go/main/App';
	export let name: string;
	export let backToLogin: () => void;
	export let onNoteClick: (note: string, index: number) => void;
	let notes: string[] = [];
	let newNote: string = '';

	async function loadNotes(): Promise<void> {
		try {
			notes = await GetNotes(name);
		} catch (error) {
			notes = [];
		}
	}

	async function addNote(): Promise<void> {
		if (newNote) {
			await AddNote(name, newNote);
			await loadNotes();
			newNote = '';
		}
	}

	loadNotes();
</script>

<div class="notes">
	<button class="back-btn" on:click={backToLogin}>Back</button>
	<h1>Notes for {name}</h1>
	<ul>
		{#each notes as note, i}
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<li on:click={() => onNoteClick(note, i)}>{note}</li>
		{/each}
	</ul>
	<textarea
		on:keydown={(event) => handleMetaEnterKeydown(event, addNote)}
		autocomplete="off"
		bind:value={newNote}
		placeholder="New note"
		class="input"
	/>
	<button class="btn" on:click={addNote}>Add Note</button>
</div>

<style>
	.notes {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.back-btn {
		position: absolute;
		left: 10px;
		width: 100px;
		height: 20px;
		border: 1px solid white;
		border-radius: 5px;
		color: white;
		background-color: transparent;
		cursor: pointer;
		margin-top: 10px;

		&:hover {
			color: black;
			background-color: white;
			transition: 0.3s linear;
		}
	}

	.notes ul {
		list-style-type: none;
		padding: 0;
	}

	.notes li {
		margin: 5px 0;
		cursor: pointer;
		border: 1px solid white;
		width: 250px;
		border-radius: 5px;
		transition: 0.3s linear;
		padding: 5px;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;

		&:hover {
			background-color: white;
			color: black;
			border-color: black;
		}
	}

	.notes .input {
		width: 250px;
		height: 100px;
		border-radius: 5px;
		border: none;
		padding: 10px;
		margin-top: 10px;
	}

	.notes .btn {
		width: 100px;
		height: 20px;
		border: 1px solid white;
		border-radius: 5px;
		color: white;
		background-color: transparent;
		cursor: pointer;
		margin-top: 10px;

		&:hover {
			color: black;
			background-color: white;
			transition: 0.3s linear;
		}
	}
</style>
