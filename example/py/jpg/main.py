import cv2
import math

img = cv2.imread(
    '/Users/wangyang/Downloads/source.png')

# cv2.imshow("ss",  img[1500:1942, 0:1920])
# cv2.waitKey()
cv2.imwrite("/Users/wangyang/Downloads/cut.png", img[1500:1942, 0:1920])
