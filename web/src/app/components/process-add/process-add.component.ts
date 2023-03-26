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
  styleUrls: [