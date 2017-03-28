 window.onload = function(){
              cc.game.onStart = function(){
                  //load resources
                  cc.LoaderScene.preload(["HelloWorld.png"], function () {
                      var MyScene = cc.Scene.extend({
                          onEnter:function () {
                              this._super();
                              var size = cc.director.getWinSize();
                              // var sprite = cc.Sprite.create("HelloWorld.png");
                              // sprite.setPosition(size.width / 2, size.height / 2);
                              // sprite.setScale(0.8);
                              // this.addChild(sprite, 0);

                              var label = cc.LabelTTF.create("Hello World", "Arial", 40);
                              label.setPosition(size.width / 2, size.height / 2);
                              this.addChild(label, 1);
                                var layer = new cc.LayerColor(cc.color(255, 0, 0, 128));
                                layer.ignoreAnchor = false;
                                layer.anchorX = 0.5;
                                layer.anchorY = 0.5;
                                layer.setContentSize(200, 200);
                                layer.x = size.width / 2;
                                layer.y = size.height / 2;
                                this.addChild(layer, 10, cc.TAG_LAYER);


                              var listener1 = cc.EventListener.create({
                                  event: cc.EventListener.TOUCH_ONE_BY_ONE,
                                  swallowTouches: true,
                                  onTouchBegan: function (touch, event) {
                                      var target = event.getCurrentTarget();                            
                                      var locationInNode = target.convertToNodeSpace(touch.getLocation());
                                      var s = target.getContentSize();
                                      var rect = cc.rect(0, 0, s.width, s.height);
                                      if (cc.rectContainsPoint(rect, locationInNode)) {
                                          cc.log("sprite began... x = " + locationInNode.x + ", y = " + locationInNode.y);
                                          target.opacity = 180;
                                          return true;
                                      }
                                      return false;
                                  },
                                  onTouchMoved: function (touch, event) {
                                      var target = event.getCurrentTarget();
                                      var delta = touch.getDelta();
                                      target.x += delta.x;
                                      target.y += delta.y;
                                  },
                                  onTouchEnded: function (touch, event) {
                                      var target = event.getCurrentTarget();
                                      cc.log("sprite onTouchesEnded.. ");
                                      target.setOpacity(255);
                                      if (target == sprite2) {
                                          containerForSprite1.setLocalZOrder(100);
                                      } else if (target == sprite1) {
                                          containerForSprite1.setLocalZOrder(0);
                                      }
                                  }
                              });
                              cc.eventManager.addListener(listener1, layer);

                          }
                      });
                      cc.director.runScene(new MyScene());
                  }, this);
              };
              cc.game.run("gameCanvas");
          };