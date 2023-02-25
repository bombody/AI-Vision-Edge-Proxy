
import { AfterViewInit, Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { AppProcess } from 'src/app/models/AppProcess';
import { EdgeService } from 'src/app/services/edge.service';
import { Terminal } from 'xterm';
import { ConfirmDialogComponent } from '../shared/confirm-dialog/confirm-dialog.component';

@Component({
  selector: 'app-app-details',
  templateUrl: './app-details.component.html',
  styleUrls: ['./app-details.component.scss']
})
export class AppDetailsComponent implements OnInit, AfterViewInit {

  private sub:Subscription;
  @ViewChild('errorTerminal') errTerminalDiv: ElementRef;
  @ViewChild('outTerminal') outTerminalDiv: ElementRef;

  term:Terminal;
  outTerm:Terminal;

  process:AppProcess = {
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
    this.edgeService.getApp(name).subscribe(proc => {

      console.log("process: ", name);

      if (proc.logs) {
        if (proc.logs.stdout) {
          let stdout = atob(proc.logs.stdout)
          this.outTerm.writeln(stdout);
        }
        if (proc.logs.stderr) {
          let stderr = atob(proc.logs.stderr)