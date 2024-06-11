<script lang="ts">
	import { handleMetaEnterKeydown } from '../utils/keydown.util';

	export let text: string;
	export let index: number;
	export let saveNote: (text: string, index: number) => void;

	function saveNoteHandler() {
		saveNote(text, index);
	}
</script>

<div
	contenteditable="true"
	class="note"
	bind:innerText={text}
	on:keydown={(event) => handleMetaEnterKeydown(event, saveNoteHandler)}
></div>
<button class="btn" on:click={() => saveNote(text, index)}>Save</button>

<style>
	.note {
		padding: 20px;
		max-height: calc(100% - 50px);
		width: calc(100% - 60px);
		text-align: left;
		overflow-y: scroll;
	}

	[contenteditable='true']:focus {
		outline: none;
	}

	.btn {
		width: 200px;
		height: 30px;
		color: white;
		border-radius: 5px;
		background-color: transparent;
		border: 1px solid white;
		font-size: 14px;
		cursor: pointer;
		transition: 0.3s linear;

		&:hover {
			background-color: white;
			color: black;
		}
	}
</style>
