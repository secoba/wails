//
//  ApDelegate.h
//  test
//
//  Created by Lea Anthony on 10/10/21.
//

#ifndef ApDelegate_h
#define ApDelegate_h

#import <Cocoa/Cocoa.h>
#import "WailsContext.h"

@interface ApDelegate : NSResponder <NSApplicationDelegate, NSTouchBarProvider>

@property bool alwaysOnTop;
@property bool startHidden;
@property (retain) NSString* singleInstanceUniqueId;
@property bool singleInstanceLockEnabled;
@property bool startFullscreen;
@property (retain) WailsWindow* mainWindow;

@end

extern void HandleOpenFile(char *);

extern void HandleSecondInstanceData(char * message);

void SendDataToFirstInstance(char * singleInstanceUniqueId, char * text);

char* GetMacOsNativeTempDir();

#endif /* ApDelegate_h */
