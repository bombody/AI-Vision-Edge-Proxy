import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { NotificationsService } from 'angular2-notifications';
import { AppProcess, PortMap } from 'src/app/models/AppProcess';
import { EdgeService } from 'src/app/services/edge.service';
import { WaitDialogComponent } from '../shared/wait-dialog/wait-dialog.component';

@Component({
  selector: 'app-app-add',
  templateUrl: './app-add.component.html',
  styleUrls: ['./app-add.component.scss']
})
export class AppAddComponent implements OnInit {

  appForm:FormGroup;
  // dockerTags:DockerTag[] = [
  //   {value: "", viewValue:"default"}
  // ]
  runtimeSelected:string = '';
  submitted:Boolean = false;
  errorMessage:string;
  loadingMessage:string;

  constructor(private _formBuilder:FormBuilder, 
    private edgeService:EdgeService, 
    private router:Router,
    private notifService:NotificationsService,
    public dialog:MatDialog) { 

    this.appForm = this._formBuilder.group({
      name: [null, [Validators.required, Validators.minLength(3)]],
      docker_user: [null],
      docker_repository: [null, Validators.required],
      docker_version: [null, Validators.required],
      env_vars: this._formBuilder.array([]),
      mount: this._formBuilder.array([]),
      port_mappings: this._formBuilder.array([]),
      runtime: [null],
    });
  }

  mounts(): FormArray {
    return this.appForm.get("mount") as FormArray;
  }

  newMount(): FormGroup {
    return this._formBuilder.group({
      name: [null, Validators.required],
      value: [null, Validators.required],
    })
  }

  addMount() {
    this.mounts().push(this.newMount());
  }

  removeMount(i:number) {
    this.mounts().removeAt(i);
  }

  portMaps(): FormArray {
    return this.appForm.get("port_mappings") as FormArray
  }

  newPortMap(): FormGroup {
    return this._formBuilder.group({
      exposed: [null, Validators.required],
      map_to: [null, Validators.required],
    })
  }

  addPortMap() {
    this.portMaps().push(this.newPortMap());
  }
   
  removePortMap(i:number) {
    this.portMaps().removeAt(i);
  }


  envVars() : FormArray {
    return this.appForm.get("env_vars") as FormArray
  }

  newEnvVar(): FormGroup {
    return this._formBuilder.group({
      name: [null, Validators.required],
      value: [null, Validators.required],
    })
  }

  addEnvVar() {
    this.envVars().push(this.newEnvVar());
  }
   
  removeEnvVar(i:number) {
    this.envVars().removeAt(i);
  }

  get f() { return this.appForm.controls; }

  ngOnInit(): void {
  }

  downloadApp(app:AppProcess, tag:string,version:string, title:string, message:string) {
    const dialogReg = this.dialog.open(WaitDialogComponent, {
      maxWidth: "400px",
      disableClose: true,
      data: {
        title: title,
        message: message
      }
    });

    console.log("inspect app: ", 