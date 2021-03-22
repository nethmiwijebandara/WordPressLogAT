import { TestBed } from '@angular/core/testing';

import { FileuploadService } from './fileupload.service';

describe('FileuploadService', () => {
  let serviceFileUpload: FileuploadService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    serviceFileUpload = TestBed.inject(FileuploadService);
  });

  it('should be created', () => {
    expect(serviceFileUpload).toBeTruthy();
  });
});
