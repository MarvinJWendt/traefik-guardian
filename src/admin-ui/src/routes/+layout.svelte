<script>
	import '../app.css';
</script>

<!--Make site smaller-->
<div class="flex justify-center shrink-0 w-screen h-screen bg-base-200">
	<div class="2xl:w-1/2 w-full 2xl:min-w-[1536px] px-8 py-8">
		<div class="card w-full h-full bg-base-100 shadow-2xl">
			<div class="card-body">
				<div class="w-full h-full">
					<slot />
				</div>
			</div>
		</div>
	</div>
</div>
