import { Component, OnInit } from '@angular/core';
import { HttpClient} from '@angular/common/http';
import {Program} from '../../models/program.model';

import { ToastService } from '../../services/toast/toast.service';

@Component({
  selector: 'app-programs-sidenavbar',
  templateUrl: './programs-sidenavbar.component.html',
  styleUrls: ['./programs-sidenavbar.component.scss']
})
export class ProgramsSidenavbarComponent implements OnInit {

  // The programList and date will be used by the HTML to display the data.
  private programsList: Program[] = [];
  private currentDate = new Date();

  constructor(private http: HttpClient, private toastService: ToastService) { }

  ngOnInit() {
    this.getPrograms();
  }

  // getPrograms gets the JSON from the backend and stores it in programs_list.
  // It also displays an error message via the toast service if one occurs.
  private getPrograms() {
    this.http.get<Program[]>('programs').subscribe(
      (response) => {
        if (response.length !== 0) {
          this.programsList = response;
        }
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }

}
