package http

var index = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>nintengo - {{.NES.ROM.GameName}}</title>

    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css">

    <!-- Optional theme -->
    <!-- <link href="//maxcdn.bootstrapcdn.com/bootswatch/3.2.0/darkly/bootstrap.min.css" rel="stylesheet"> -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap-theme.min.css">
  </head>
  <body>
    <div class='container'>
      <div class='row'>
	<div class="page-header">
	  <h1>nintengo <small>{{.NES.ROM.GameName}}</small></h1>
	</div>

	<div class='col-md-6'>
	  <table class='table table-striped'>

	    <thead><tr><td><strong>CPU Variable</strong></td><td><strong>Value</strong></td></tr></thead>
	    <tbody>
	      <tr><td><kbd>A</kbd></td>   <td><code>{{printf "$%02x" .NES.CPU.M6502.Registers.A}}</code></td></tr>
	      <tr><td><kbd>X</kbd></td>   <td><code>{{printf "$%02x" .NES.CPU.M6502.Registers.X}}</code></td></tr>
	      <tr><td><kbd>Y</kbd></td>   <td><code>{{printf "$%02x" .NES.CPU.M6502.Registers.Y}}</code></td></tr>
	      <tr><td><kbd>P</kbd></td>   <td><code>{{printf "$%02x" .NES.CPU.M6502.Registers.P}}</code></td></tr>
	      <tr><td><kbd>SP</kbd></td>  <td><code>{{printf "$%04x" .NES.CPU.M6502.Registers.SP}}</code></td></tr>
	      <tr><td><kbd>PC</kbd></td>  <td><code>{{printf "$%04x" .NES.CPU.M6502.Registers.PC}}</code></td></tr>
	      <tr><td><kbd>NMI</kbd></td> <td><code>{{.NES.CPU.M6502.Nmi}}</code></td> </tr>
	      <tr><td><kbd>IRQ</kbd></td> <td><code>{{.NES.CPU.M6502.Irq}}</code></td> </tr>
	      <tr><td><kbd>RST</kbd></td> <td><code>{{.NES.CPU.M6502.Rst}}</code></td> </tr>
	    </tbody>

	    <thead><tr><td><strong>DMA Variable</strong></td><td><strong>Value</strong></td></tr></thead>
	    <tbody>
	      <tr><td><kbd>Pending</kbd></td> <td><code>{{printf "$%04x" .NES.CPU.DMA.Pending}}</code></td> </tr>
	    </tbody>

	    <thead><tr><td><strong>APU Variable</strong></td><td><strong>Value</strong></td></tr></thead>
	    <tbody>
	      <tr><td><kbd>Control</kbd></td>   <td><code>{{printf "$%02x" .NES.CPU.APU.Registers.Control}}</code></td></tr>
	      <tr><td><kbd>Status</kbd></td>    <td><code>{{printf "$%02x" .NES.CPU.APU.Registers.Status}}</code></td></tr>
	    </tbody>

	    <thead><tr><td><strong>PPU Variable</strong></td><td><strong>Value</strong></td></tr></thead>
	    <tbody>
	      <tr><td><kbd>Frame</kbd></td>    <td><code>{{.NES.PPU.Frame}}</code></td></tr>
	      <tr><td><kbd>Scanline</kbd></td> <td><code>{{.NES.PPU.Scanline}}</code></td></tr>
	      <tr><td><kbd>Cycle</kbd></td>    <td><code>{{.NES.PPU.Cycle}}</code></td></tr>

	      <tr><td><kbd>Controller</kbd></td> <td><code>{{printf "$%02x" .NES.PPU.Registers.Controller}}</code></td></tr>
	      <tr><td><kbd>Mask</kbd></td>       <td><code>{{printf "$%02x" .NES.PPU.Registers.Mask}}</code></td></tr>
	      <tr><td><kbd>Status</kbd></td>     <td><code>{{printf "$%02x" .NES.PPU.Registers.Status}}</code></td></tr>
	      <tr><td><kbd>OAMAddress</kbd></td> <td><code>{{printf "$%02x" .NES.PPU.Registers.OAMAddress}}</code></td></tr>
	      <tr><td><kbd>Scroll</kbd></td>     <td><code>{{printf "$%04x" .NES.PPU.Registers.Scroll}}</code></td></tr>
	      <tr><td><kbd>Address</kbd></td>    <td><code>{{printf "$%04x" .NES.PPU.Registers.Address}}</code></td></tr>
	      <tr><td><kbd>Data</kbd></td>       <td><code>{{printf "$%02x" .NES.PPU.Registers.Data}}</code></td></tr>

	      <tr><td><kbd>Latch</kbd></td>        <td><code>{{.NES.PPU.Latch}}</code></td></tr>
	      <tr><td><kbd>LatchAddress</kbd></td> <td><code>{{printf "$%04x" .NES.PPU.LatchAddress}}</code></td></tr>
	      <tr><td><kbd>LatchValue</kbd></td>   <td><code>{{printf "$%02x" .NES.PPU.LatchValue}}</code></td></tr>

	      <tr><td><kbd>AddressLine</kbd></td>    <td><code>{{printf "$%04x" .NES.PPU.AddressLine}}</code></td></tr>
	      <tr><td><kbd>PatternAddress</kbd></td> <td><code>{{printf "$%04x" .NES.PPU.PatternAddress}}</code></td></tr>

	      <tr><td><kbd>AttributeNext</kbd></td>  <td><code>{{printf "$%02x" .NES.PPU.AttributeNext}}</code></td></tr>
	      <tr><td><kbd>AttributeLatch</kbd></td> <td><code>{{printf "$%02x" .NES.PPU.AttributeLatch}}</code></td></tr>
	      <tr><td><kbd>Attributes</kbd></td>     <td><code>{{printf "$%04x" .NES.PPU.Attributes}}</code></td></tr>

	      <tr><td><kbd>TilesLatch</kbd></td> <td><code>{{printf "$%04x" .NES.PPU.TilesLatch}}</code></td></tr>
	      <tr><td><kbd>TilesLow</kbd></td>   <td><code>{{printf "$%04x" .NES.PPU.TilesLow}}</code></td></tr>
	      <tr><td><kbd>TilesHigh</kbd></td>  <td><code>{{printf "$%04x" .NES.PPU.TilesHigh}}</code></td></tr>

	    <thead><tr><td><strong>OAM Variable</strong></td><td><strong>Value</strong></td></tr></thead>
	    <tbody>
	      <tr><td><kbd>Address</kbd></td>  <td><code>{{printf "$%04x" .NES.PPU.OAM.Address}}</code></td></tr>
	      <tr><td><kbd>Latch</kbd></td>    <td><code>{{printf "$%02x" .NES.PPU.OAM.Latch}}</code></td></tr>
	      <tr><td><kbd>SpriteZeroInBuffer</kbd></td>    <td><code>{{.NES.PPU.OAM.SpriteZeroInBuffer}}</code></td></tr>
	    </tbody>

	  </table>

	  <h4>CPU Memory</h4>
	  <pre style='font-size: 11px' class='pre-scrollable'>{{.CPUMemory}}</pre>

	  <h4>PPU Memory</h4>
	  <pre style='font-size: 11px' class='pre-scrollable'>{{.PPUMemory}}</pre>

	</div>
	<div class='col-md-6'>

	  <table class='table table-striped'>
            {{range $i, $s := .NES.PPU.Sprites}}
	      <thead><tr><td><strong>PPU Sprite {{$i}} Variable</strong></td><td><strong>Value</strong></td></tr></thead>
	      <tbody>

        	<tr><td><kbd>TileLow </kbd></td>   <td><code>{{printf "$%0x2" $s.TileLow}}</code></td></tr>
        	<tr><td><kbd>TileHigh </kbd></td>  <td><code>{{printf "$%0x2" $s.TileHigh}}</code></td></tr>
        	<tr><td><kbd>Sprite </kbd></td>    <td><code>{{printf "$%08x" $s.Sprite}}</code></td></tr>
        	<tr><td><kbd>XPosition </kbd></td> <td><code>{{printf "$%02x" $s.XPosition}}</code></td></tr>
        	<tr><td><kbd>Address </kbd></td>   <td><code>{{printf "$%04x" $s.Address}}</code></td></tr>
        	<tr><td><kbd>Priority </kbd></td>  <td><code>{{printf "$%02x" $s.Priority}}</code></td></tr>
            {{end}}
	      </tbody>
	  </table>

	</div>
      </div>
      <div class='row'>

	<div class='col-md-8'>
	  <h4>PPU Pattern Tables</h4>
	</div>

	<div class='col-md-6'>
          <img alt='left'  style='width:100%;' src='data:image/png;base64,{{.PTLeft}}'  class='img-thumbnail img-responsive' />
	</div>

	<div class='col-md-6'>
          <img alt='right' style='width:100%;' src='data:image/png;base64,{{.PTRight}}' class='img-thumbnail img-responsive' />
	</div>

      </div>
    </div>

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
    <!-- Latest compiled and minified JavaScript -->
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
  </body>
</html>
`