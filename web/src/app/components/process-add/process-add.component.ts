import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { EdgeService } from 'src/app/services/edge.service';
import { StreamProcess } from 'src/app/models/StreamProcess';
import { Router } from '@angular/router';
import { NotificationsService } from 'angular2-notifications';

interface DockerTag {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'app-process-add',
  templateUrl: './process-add.component.html',
  styleUrls: ['./process-add.component.scss']
})
export class ProcessAddComponent implements OnInit {

  rtspForm:FormGroup;
  dockerTags:DockerTag[] = [
    {value: "", viewValue:"default"}
  ]
  tagSelected:string = '';
  submitted:Boolean = false;
  errorMessage:string;

  constructor(private _formBuilder:FormBuilder, private edgeService:EdgeService, private router:Router, private notifService:NotificationsService) {
    this.rtspForm = this._formBuilder.group({
      name: [null, [Validators.required, Validators.minLength(4)]],
      rtsp_endpoint: [null, [Validators.required]],
      rtmp_endpoint: [null],
    });
   }

  ngOnInit(): void {

    let data = history.state.data;
    if (data) {
      if (data.rtsp_endpoint) {
        this.rtspForm.get('rtsp_endpoint').setValue(data.rtsp_endpoi