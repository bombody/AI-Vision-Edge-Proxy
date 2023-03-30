
import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { StreamProcess } from 'src/app/models/StreamProcess';
import { EdgeService } from 'src/app/services/edge.service';
import { Terminal } from 'xterm';
import { ConfirmDialogComponent } from '../shared/confirm-dialog/confirm-dialog.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-process-details',
  templateUrl: './process-details.component.html',
  styleUrls: ['./process-details.component.scss'],
})
export class ProcessDetailsComponent implements OnInit, AfterViewInit {

  private sub:Subscription;
  @ViewChild('errorTerminal') errTerminalDiv: ElementRef;
  @ViewChild('outTerminal') outTerminalDiv: ElementRef;

  term:Terminal;
  outTerm:Terminal;

  process:StreamProcess = {
  };

  constructor(private route:ActivatedRoute, private edgeService:EdgeService, public dialog:MatDialog, private router:Router) {

  }
  ngAfterViewInit(): void {
    this.term = new Terminal({
      convertEol: true,
      logLevel: "off",
      rendererType: "canvas",
      theme: {foreground: '#ff0000'},
    });
    this.term.open(this.errTerminalDiv.nativeElement);

    this.outTerm = new Terminal({
      allowTransparency: true,
      convertEol: true,
      logLevel: "info"
    })
    this.outTerm.open(this.outTerminalDiv.nativeElement);
  }

  ngOnInit(): void {


    this.sub = this.route.params.subscribe(params => {
      this.process.name = params['name'];
      this.loadProcess(this.process.name);
    });
  }

  loadProcess(name:string) {
    this.edgeService.getProcess(name).subscribe(proc => {
      if (proc.logs) {
        if (proc.logs.stdout) {
          let stdout = atob(proc.logs.stdout)
          this.outTerm.writeln(stdout);
        }
        if (proc.logs.stderr) {
          let stderr = atob(proc.logs.stderr)
          console.log(stderr);
          this.term.writeln("\x1B[1;3;31m=====ERROR LOGS=====\x1B[0m");
          this.term.writeln(stderr);
        }
      }
      console.log(proc);
      this.process = proc;
    }, error => {
      console.error(error);
    })
  }

  delete(process:StreamProcess) {
    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      maxWidth: "400px",
      data: {
          title: "Are you sure?",
          message: "You are about to delete RTSP camera stream."}
      });
  
      // listen to response
      dialogRef.afterClosed().subscribe(dialogResult => {
        // if user pressed yes dialogResult will be true, 
        // if he pressed no - it will be false
        if (dialogResult) {
          console.log("delete: ", dialogResult);
          this.edgeService.stopProcess(process.name).subscribe(res => {
            console.log("delete success: ", res);
            this.router.navigate(['/local/processes']);
          }, error => {
            console.error(error);
          });
        }
      });
  }

}