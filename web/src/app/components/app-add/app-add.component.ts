import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { NotificationsService } from 'angular2-notifications';
import { AppProcess, PortMap } from 'src/app/models/AppProcess';
import { EdgeService }