<div class="schematic-buttons">
	<button class="schematic-button-pause schematic-button" title="Pause/resume the simulation">
		<!-- only one of these two images is visible at a time -->
		<img src="/playground/resources/codicon/debug-pause.svg" class="button-img-pause"/>
		<img src="/playground/resources/codicon/play.svg" class="button-img-play"/>
	</button>
</div>
<svg class="schematic" tabindex="0">
	<g class="schematic-wrapper" style="transform: translate(50%, 50%)">
		<g class="schematic-parts"></g>
		<g class="schematic-wires"></g>
	</g>
</svg>
<div class="card-header">
	<ul class="nav nav-tabs card-header-tabs" role="tablist">
		<li class="nav-item" role="presentation">
			<button class="nav-link active panel-tab-terminal" id="simulator-tab-terminal" data-bs-toggle="tab" data-bs-target="#simulator-panel-terminal" type="button" role="tab" aria-controls="simulator-panel-terminal" aria-selected="true">Terminal</button>
		</li>
		<li class="nav-item" role="presentation">
			<button class="nav-link" id="simulator-properties-tab" data-bs-toggle="tab" data-bs-target="#simulator-panel-properties" type="button" role="tab" aria-controls="simulator-panel-properties" aria-selected="false">Properties</button>
		</li>
		<li class="nav-item" role="presentation">
			<button class="nav-link" id="simulator-power-tab" data-bs-toggle="tab" data-bs-target="#simulator-panel-power" type="button" role="tab" aria-controls="simulator-panel-power" aria-selected="false">Power</button>
		</li>
		<li class="nav-item" role="presentation">
			<button class="nav-link" id="simulator-add-tab" data-bs-toggle="tab" data-bs-target="#simulator-panel-add" type="button" role="tab" aria-controls="simulator-panel-add" aria-selected="false">Add</button>
		</li>
	</ul>
</div>
<div class="tab-content">
	<div class="tab-pane active terminal-box" id="simulator-panel-terminal" role="tabpanel" aria-labelledby="simulator-tab-terminal">
		<div class="terminal" tabindex="0"></div>
	</div>
	<div class="tab-pane panel-properties content" id="simulator-panel-properties" role="tabpanel" aria-labelledby="simulator-properties-tab">
		<div class="content" tabindex="0"></div>
	</div>
	<div class="tab-pane panel-power content" id="simulator-panel-power" role="tabpanel" aria-labelledby="simulator-power-tab">
		<div class="content" tabindex="0">
			<div class="power-table">Loading...</div>
			<p class="mb-0 mt-3"><strong>Note:</strong> these numbers are estimates, based on datasheets and measurements. They don't include everything and may be wrong.</p>
		</div>
	</div>
	<div class="tab-pane panel-add" id="simulator-panel-add" role="tabpanel" aria-labelledby="simulator-add-tab">
			<div class="content" tabindex="0">Loading...</div>
	</div>
</div>
<div class="schematic-tooltip"></div>
<div class="templates d-none">
	<button class="panel-add-button btn btn-primary btn-sm">Add</button>
	<select class="panel-add-select form-select form-select-sm"></select>
</div>
