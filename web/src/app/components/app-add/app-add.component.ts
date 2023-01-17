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
  //   {value: "", viewValu